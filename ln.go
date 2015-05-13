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
import flag "github.com/ogier/pflag"

func main() {
    sFlag := flag.BoolP("symbolic", "s", false, "symlink mode")
    flag.Parse();
    if(len(flag.Args())>1) {
        if(*sFlag) {
            os.Symlink(flag.Arg(0) , flag.Arg(1))
        } else {
            os.Link(flag.Arg(0) , flag.Arg(1))
        }
    }
}
