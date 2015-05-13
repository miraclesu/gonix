/*
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License version 3 as
	published by the Free Software Foundation.
	
	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	
	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import "os"
import "io"

func main() {
	//sf: source file, df: destination file
	sf, _ := os.Open(os.Args[1])
	sfinfo, _ := os.Stat(os.Args[1])
	sfmode := sfinfo.Mode()
	df, _ := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, sfmode)
	io.Copy(df, sf)
	sf.Close()
	df.Close()
}
