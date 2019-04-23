/*********************************************
                   _ooOoo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  O\  =  /O
               ____/`---'\____
             .'  \\|     |//  `.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||-  \
           |   | \\\  -  /// |   |
           | \_|  ''\---/''  |   |
           \  .-\__  `-`  ___/-. /
         ___`. .'  /--.--\  `. . __
      ."" '<  `.___\_<|>_/___.'  >'"".
     | | :  `- \`.;`\ _ /`;.`/ - ` : | |
     \  \ `-.   \_ __\ /__ _/   .-` /  /
======`-.____`-.___\_____/___.-`____.-'======
                   `=---='

^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
           佛祖保佑       永无BUG
           心外无法       法外无心
           三宝弟子       飞猪宏愿
*********************************************/

package main

import (
	"fmt"
	"strconv"

	"github.com/alash3al/go-color"
)

func main() {
	fmt.Println(color.MagentaString(serverBrand))
	fmt.Printf("  服务版本: %s \n", color.GreenString(serverVersion))
	fmt.Printf("  SQL 驱动: %s \n", color.GreenString(*flagDBDriver))
	fmt.Printf("  连接字串: %s \n", color.GreenString(*flagDBDSN))
	fmt.Printf("  工作线程: %s \n", color.GreenString(strconv.Itoa(*flagWorkers)))
	fmt.Printf("  监听地址: %s \n", color.GreenString(*flagRESTListenAddr))

	if *flagRedisURL == "" {
		fmt.Printf("  Redis 缓存: %s \n", color.RedString("<未配置>"))
	} else {
		fmt.Printf("  Redis 缓存: %s \n", color.GreenString(*flagRedisURL))
	}

	if *flagRSAPrivkey == "" || *flagJWTSecret == "" {
		fmt.Printf("  JWT RSA私钥: %s \n", color.RedString("<未配置>"))
		fmt.Printf("  JWT 安全令牌: %s \n", color.RedString("<未配置>"))
		fmt.Printf("  JWT 令牌期限: %s秒 \n", color.GreenString(strconv.Itoa(*flagJWTExpires)))
	} else {
		fmt.Printf("  JWT RSA私钥: %s \n", color.GreenString(*flagRSAPrivkey))
		fmt.Printf("  JWT 安全令牌: %s \n", color.GreenString(*flagJWTSecret))
		fmt.Printf("  JWT 令牌期限: %s秒 \n", color.GreenString(strconv.Itoa(*flagJWTExpires)))
	}

	if *flagJWTSecret == "" {
	} else {
	}

	err := make(chan error)

	go (func() {
		err <- initRestfulServer()
	})()

	if err := <-err; err != nil {
		color.Red(err.Error())
	}
}
