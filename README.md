# go-jq
<img width="694" alt="image" src="https://user-images.githubusercontent.com/18223385/222345227-859084de-511c-40dc-9e8e-f4ed964409c4.png">

# What Features
- Resolve the exception such as parse error: Invalid uXXXX uXXXX in jq(https://github.com/stedolan/jq/issues/2543)
- Year-on-year jq, supporting the setting of golang format

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
```

## What format
is golang format
```
"%v:%v"
```
