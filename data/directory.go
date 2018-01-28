/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data

type Directory struct {
	ID int
	ParentID int
	Title string
	Path string
}

func (d *Directory) SubDirectories() ([]Directory, error) {
	return DataBase.SubDirectories(d.ID)
}

func (d *Directory) Save() error {
	return DataBase.SaveDirectory(d)
}

func (d *Directory) Delete() error {
	return DataBase.DeleteDirectory(d)
}

func (d *Directory) Load() error {
	return DataBase.LoadDirectory(d)
}