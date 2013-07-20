/*
 * SimpleGPIO.go
 *
 * Translated into Go by Frank Davidson, ffdavidson@gmail.com
 * Specific to the Beaglebone Black
 *
 * Based on SimpleGPIO.cpp
 *
 * Modifications by Derek Molloy, School of Electronic Engineering, DCU
 * www.derekmolloy.ie
 * Almost entirely based on Software by RidgeRun:
 *
 * Copyright (c) 2011, RidgeRun
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *    This product includes software developed by the RidgeRun.
 * 4. Neither the name of the RidgeRun nor the
 *    names of its contributors may be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY RIDGERUN ''AS IS'' AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL RIDGERUN BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package SimpleGPIO

import (
	"os"
	"strconv"
	"syscall"
)

const SYSFS_GPIO_DIR, POLL_TIMEOUT = "/sys/class/gpio", (3 * 1000) // 3 seconds

type PIN_DIRECTION int

const (
	INPUT_PIN PIN_DIRECTION = iota
	OUTPUT_PIN
)

type PIN_VALUE int

const (
	LOW PIN_VALUE = iota
	HIGH
)

func GPIOExport(gpio int) {

	gpioBuf := []byte(strconv.Itoa(gpio)) // Put the gpio int into a byte array as char data

	buf, err := os.OpenFile(SYSFS_GPIO_DIR+"/export", syscall.O_WRONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	_, err = buf.Write(gpioBuf)

	if err != nil {
		panic(err)
	}

	return
}

func GPIOUnexport(gpio int) {

	gpioBuf := []byte(strconv.Itoa(gpio)) // Put the gpio int into a byte array as char data

	buf, err := os.OpenFile(SYSFS_GPIO_DIR+"/unexport", syscall.O_WRONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	_, err = buf.Write(gpioBuf)

	if err != nil {
		panic(err)
	}

	return
}

func GPIOSetDirection(gpio int, pinDir PIN_DIRECTION) {

	in := []byte("in")
	out := []byte("out")

	buf, err := os.OpenFile(SYSFS_GPIO_DIR+"/gpio"+strconv.Itoa(gpio)+"/direction", syscall.O_WRONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	if pinDir == INPUT_PIN {
		_, err = buf.Write(in)
	} else {
		_, err = buf.Write(out)
	}

	return
}

func GPIOSetValue(gpio int, value PIN_VALUE) {

	valueBuf := []byte(strconv.Itoa(int(value))) // Put the gpio int into a byte array as char data

	buf, err := os.OpenFile(SYSFS_GPIO_DIR+"/gpio"+strconv.Itoa(gpio)+"/value", syscall.O_WRONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	_, err = buf.Write(valueBuf)

	if err != nil {
		panic(err)
	}

	return
}

func GPIOGetValue(gpio int) PIN_VALUE {

	value := make([]byte, 1)

	buf, err := os.OpenFile(SYSFS_GPIO_DIR+"/gpio"+strconv.Itoa(gpio)+"/value", syscall.O_RDONLY, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	_, err = buf.Read(value)

	if err != nil {
		panic(err)
	}

	if value[0] == '1' {
		return HIGH
	} else {
		return LOW
	}

}
