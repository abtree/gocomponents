package main

import "config/decode"

/*
关键字：
	1 .key [列序号,列序号] 运行配置多列为key，默认为第一列
	2 .multi 运行配置repeated的key
*/
/*
	1 数组配置，支持分列配或合并在一起配置(且可混用)
		->分列 []uint32:Param	:Param  :Param ==>	1 \t 2 \t 3
		->合并 []uint32:_Param	==>	1_2_3
	2 允许struct合并配置
		CPriceItem:_Award ===> 1_EC001_100	合并配置可能有些字段不会填值 从前往后填 注意会忽略多出来的struct字段
	3 允许子pb结构的生成以解析复杂结构数据 ｛｝表示需要生成一个子结构
		如解析: `10:1_2_3_4|30:5_6_7_8`
		结构生成格式： `[{":a":"uint32","_b":["uint32"]}]`
		生成的子结构：
		message CfgXXX{
			uint32 A = 1;
	 		repeated uint32 B = 2;
		}
		repeated CfgXXX xxx = 1;

		如解析: `10:4:3|30:5:8`
		结构生成格式： `[{":a":["uint32"]}]`
		生成的子结构：
		message CfgXXX{
	 		repeated uint32 A = 1;
		}
		repeated CfgXXX xxx = 1;

		如解析: `10:4_3_5:3|30:5_7_9:8`
		结构生成格式： `[{":a":"uint32","_b":["uint32"],":c":"uint32"}]`
		生成的子结构：
		message CfgXXX{
			uint32 A = 1;
	 		repeated uint32 B = 2;
			int32 C = 3;
		}
		repeated CfgXXX xxx = 1;
	4 json解析的修改，主要是对数组的支持
	source = `{
		"a":1,
		"b":[1,2,3],
		"c":{
			"d":4.2,
			"e":5.4
		},
		"f":[{
				"g":"str1",
				"h":"str2"
			},{
				"g":"str3",
				"h":"str4"
			}]
		}`
	5 取消了对非pb解析的支持 和 单文件解析为多个 pb 的支持
	6 取消了ini文件的变量替换功能
	7 取消了对非pb结构的struct解析的支持（只能用pb解析）
	8 .multi 关键字可以和 .key 关键字同时使用
	9 .multi 可以自己生成子结构
	10 允许配置额外的字段（pb中会解析，但配置中不用）
	11	ini文件的值，允许配置为json(客户端估计不干)
		;值为json的数据
		JsonDat={"a":123,"b":"xxx"}
		;值为json数组(reflect解析不出来，会被当做[]interface{}解析)
		;JArrDat=[{"c":456,"d":"xxx"},{"c":789,"d":"yyy"}]
*/
/*
    活动配置会根据活动名称 和 子活动序列
	解析为map
*/

func main() {
	//src.Run()
	decode.Run()
}
