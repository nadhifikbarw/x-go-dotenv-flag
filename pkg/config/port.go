// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Nadhif Ikbar Wibowo

package config

import (
	"errors"
	"strconv"
)

type Port int

func (p *Port) Decode(value string) error {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if intVal < 0 || intVal > 65535 {
		return errors.New("port must be between 1-65536")
	}

	*p = Port(intVal)

	return nil
}

func (p Port) String() string {
	return strconv.Itoa(int(p))
}

func (p Port) Int() int {
	return int(p)
}
