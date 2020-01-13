package pov

type Node struct {
	name   string
	parent *Node
}

type Graph struct {
	nodes []*Node
}

func New() *Graph {
	return &Graph{nil}
}

func (g *Graph) AddNode(nodeLabel string) {
	g.nodes = append(g.nodes, &Node{nodeLabel, nil})
}

func (g *Graph) AddArc(fr, to string) {
	var p *Node

	for _, v := range g.nodes {
		if v.name == fr {
			p = v
			break
		}
	}

	if p == nil {
		p = &Node{fr, nil}
		g.nodes = append(g.nodes, p)
	}

	for _, v := range g.nodes {
		if v.name == to {
			v.parent = p
			break
		}
	}
}

func (g *Graph) ArcList() (res []string) {
	for _, v := range g.nodes {
		if v.parent != nil {
			res = append(res, v.parent.name+" -> "+v.name)
		}
	}
	return
}

func (g *Graph) copy() *Graph {
	newG := New()

	newG.nodes = make([]*Node, 0, len(g.nodes))

	for _, v := range g.nodes {
		newG.AddNode(v.name)
	}

	for _, v := range g.nodes {
		if v.parent != nil {
			newG.AddArc(v.parent.name, v.name)
		}
	}

	return newG
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {

	// g = g.copy()

	var newR *Node

	for _, v := range g.nodes {
		if v.name == newRoot {
			newR = v
		}
	}

	var ns []*Node

	for newR != nil {
		ns, newR = append(ns, newR), newR.parent
	}

	for i := len(ns) - 1; i >= 0; i-- {
		if i == 0 {
			ns[i].parent = nil
		} else {
			ns[i].parent = ns[i-1]
		}
	}

	return g
}
