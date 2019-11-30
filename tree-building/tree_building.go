/**
reference https://exercism.io/tracks/go/exercises/tree-building/solutions/db3eeec81b204483824c3f1f26499987
nice solution
**/

package tree
import (
    "fmt"
    // "sort"
)

type Node struct {
    ID int
    Children []*Node
}

type Record struct {
    ID int
    Parent int
}


func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}
	// use sort.Sort() will spend more time
	// Less() do the extra work 
    for i, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds record id %d", r.ID)
		}
		if r.ID != i {
			records[i], records[r.ID] = records[r.ID], records[i]
		}
	}

	// //check continuous
	// maxId := records[len(records)- 1].ID

	// if maxId != len(records) - 1 {
	// 	return nil, fmt.Errorf("non-continuous")
	// }

	cache := make([]Node, len(records)) 
    
    for i, r := range records {

		//check duplicate node
		if i != r.ID {
			return nil, fmt.Errorf("%v duplicate node", r)
		}
		
		if i == 0 {
			//check id == parent
			if r.Parent != 0 {
				return nil, fmt.Errorf("%v id == parent", r)
			}
			cache[i].ID = i
			continue
		}

		//check higher id parent of lower id
		if i <= r.Parent {
			return nil, fmt.Errorf("%v higher id parent of lower id", r)
		}

        cache[r.Parent].Children = append(cache[r.Parent].Children, &cache[i])
        cache[i].ID = i
        
    }
    
    return &cache[0], nil
}
