package utils

import "runtime"

/**

 * Created by cxnky on 24/04/2019 at 16:34
 * utils
 * https://github.com/cxnky/

**/

// IsInProduction checks whether the bot is running in a Windows environment (development) or a Linux one (production)
func IsInProduction() bool {
	return runtime.GOOS == "linux"
}
