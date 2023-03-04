# go-jq
<img width="694" alt="image" src="https://user-images.githubusercontent.com/18223385/222345227-859084de-511c-40dc-9e8e-f4ed964409c4.png">
<img width="715" alt="image" src="https://user-images.githubusercontent.com/18223385/222897726-536a9a4b-d631-45d2-8000-13033e84e6fd.png">


# What Features
- Resolve the exception such as parse error: Invalid uXXXX uXXXX in jq(https://github.com/stedolan/jq/issues/2543)
- Year-on-year jq, supporting the setting of golang format
- support parse nmap / masscan result xml，完美解析 xml结果


# How use
```
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
```
query
```
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Deer Hunter"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig"

friends.#(last=="Murphy").first    >> "Dale"
friends.#(last=="Murphy")#.first   >> ["Dale","Jane"]
friends.#(age>45)#.last            >> ["Craig","Murphy"]
friends.#(first%"D*").last         >> "Murphy"
friends.#(first!%"D*").last        >> "Craig"
friends.#(nets.#(=="fb"))#.first   >> ["Dale","Roger"]
```

# How build and run
```
go build -o jq main.go
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "%v:%v" "ip_str" "port"
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "ip_str" "port"
./jq $HOME/MyWork/scan4all/atckData/china_chengdu.json "ip_str" "port"
# 完美解析 nmap xml结果
./jq $HOME/MyWork/scan4all/atckData/x01/a0988c54b5a57d258a43a0a95f54e5975aaec96e.xml "%v:%v"  "nmaprun.host.#.address.addr"  "nmaprun.host.#.ports.port.#.portid"
./jq $HOME/MyWork/scan4all/atckData/x01/a0988c54b5a57d258a43a0a95f54e5975aaec96e.xml "%v:%v %v %v"  "nmaprun.host.#.address.addr"  "nmaprun.host.#.ports.port.#.portid" "nmaprun.host.#.ports.port.#.service.name" "nmaprun.host.#.ports.port.#.state.state"|grep ' open' 
```

## What format
is golang format
```
"%v:%v"
```
