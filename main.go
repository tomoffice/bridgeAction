package main

import "fmt"

type Drive interface {
	Sound()
}

type SSD struct{}

func (s *SSD) Sound() {
	fmt.Println("SSD 運轉聲...")
}

type HD struct{}

func (h *HD) Sound() {
	fmt.Println("HD 運轉聲...")
}

type Memory interface {
	Allocate()
	RW()
}

type DDR4 struct{}

func (d *DDR4) Allocate() {
	fmt.Println("配置 DDR4 記憶體...")
}

func (d *DDR4) RW() {
	fmt.Println("DDR4 讀寫中...")
}

type DDR5 struct{}

func (d *DDR5) Allocate() {
	fmt.Println("配置 DDR5 記憶體...")
}

func (d *DDR5) RW() {
	fmt.Println("DDR5 讀寫中...")
}

type Graphic interface {
	Image()
	Game()
}

type BuiltIn struct{}

func (b *BuiltIn) Image() {
	fmt.Println("使用內建顯示卡播放影像...")
}

func (b *BuiltIn) Game() {
	fmt.Println("使用內建顯示卡遊玩遊戲...")
}

type Discrete struct{}

func (d *Discrete) Image() {
	fmt.Println("使用獨立顯示卡播放影像...")
}

func (d *Discrete) Game() {
	fmt.Println("使用獨立顯示卡遊玩遊戲...")
}

type Bridge interface {
	On()
	Off()
}

type Computer struct {
	hd    Drive
	mem   Memory
	graph Graphic
}

func (c *Computer) On() {
	fmt.Println("開機中...")
	fmt.Println("使用硬碟：")
	c.hd.Sound()
	fmt.Println("使用記憶體：")
	c.mem.Allocate()
	c.mem.RW()
	fmt.Println("使用顯示卡：")
	c.graph.Image()
	c.graph.Game()
}

func (c *Computer) Off() {
	fmt.Println("關機中...")
	fmt.Println("使用記憶體：")
	c.mem.RW()
	fmt.Println("使用硬碟：")
	c.hd.Sound()
	fmt.Println("使用顯示卡：")
	c.graph.Image()
}
func NewComputer(hd Drive, mem Memory, graph Graphic) Bridge {
	return &Computer{hd: hd, mem: mem, graph: graph}
}
func main() {
	hd := &SSD{}
	mem := &DDR4{}
	graph := &BuiltIn{}
	computer := NewComputer(hd, mem, graph)
	computer.On()
	computer.Off()
}
