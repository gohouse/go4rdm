package config

type RedisCmd struct {
	CmdName string
	Summary string `json:"summary"`
	Complexity string `json:"complexity"`
	Since string `json:"since"`
	Group string `json:"group"`
	Arguments interface{} `json:"arguments"`
}

type Arguments struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Optional string  `json:"optional"`
	Enum string  `json:"enum"`
}

type Cmds struct {
	Cmd []string
	Summary []string
}

//func LoadRedisCmd(filepath string)  {
//	f, err := ioutil.ReadFile(filepath)
//	if err!=nil {
//		panic(err.Error())
//	}
//	var js = make(map[string]RedisCmd)
//
//	err = json.Unmarshal(f, &js)
//	if err!=nil {
//		panic(err.Error())
//	}
//
//	var group = make(map[string][]RedisCmd)
//	for k,v := range js {
//		// 保存 string
//		// 保存 hash
//		// 保存 set
//		// 保存 zset
//		// 保存 list
//		if v.Group == "string" ||
//			v.Group == "set" ||
//			v.Group == "hash" ||
//			v.Group == "sorted_set" ||
//			v.Group == "pubsub" ||
//			v.Group == "list" ||
//			v.Group == "geo" ||
//			v.Group == "hyperloglog" {
//			v.CmdName = k
//			group[v.Group]  = append(group[v.Group], v)
//		}
//		//group[v.Group] = v.Group
//		//fmt.Println(v.Group)
//	}
//
//	//fmt.Println("len:", len(group))
//	for k,_:= range group {
//		fmt.Println(k)
//	}
//	indent, _ := json.MarshalIndent(group, "", "  ")
//
//	file.FilePutContents("xx.json", indent)
//	fmt.Printf("%s",indent)
//}
