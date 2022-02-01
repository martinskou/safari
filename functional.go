package safari



func Map[N any, M any](lst []N, f func(N) M) []M {
	vsm := make([]M, len(lst))
	for i, v := range lst {
		vsm[i] = f(v)
	}
	return vsm
}

/* Functional filter function */
func Filter[N any](vs []N, f func(N) bool) []N {
    filtered := make([]N, 0)
    for _, v := range vs {
        if f(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}

/* Functional filter function for maps */
func FilterMap[N any](vs map[string]N, f func(N) bool) map[string]N {
    filtered := make(map[string]N, 0)
    for k, v := range vs {
        if f(v) {
            filtered[k]=v
        }
    }
    return filtered
}

/* Create a map from a slice, with key defined by keyfunc */
func SliceToMap[N any](lst []N, keyfunc func(N)string ) map[string]N {
	mp:=make(map[string]N)
	for i:=0;i<len(lst);i++ {
		mp[keyfunc(lst[i])]=lst[i]
	}
	return mp
} 

/* Return all keys from a map */
func MapKeys[N any](mp map[string]N) []string {
	lst:=[]string{}
	for k,_:=range mp {
		lst = append(lst, k)
	}	
	return lst
}

/* Return all values from a map */
func MapValues[N any](mp map[string]N) []N {
	lst:=[]N{}
	for _,v:=range mp {
		lst = append(lst, v)
	}	
	return lst
}
