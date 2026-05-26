package main

import "fmt"

type ll struct {
	val  int
	prox *ll
}

func add_next(val int, head *ll) {
	if head == nil {
		return
	}
	analise := head
	for {
		if analise.prox == nil {
			break
		}
		analise = analise.prox
	}
	proximo := ll{val: val, prox: nil}
	analise.prox = &proximo
}

func show_list(head *ll) {
	if head == nil {
		return
	}
	analise := head
	for {
		fmt.Printf("%v -> ", analise.val)
		if analise.prox == nil {
			break
		}
		analise = analise.prox
	}
}

func inverter(head *ll) *ll {
	var n *ll = nil
	var resto *ll = nil

	for {
		if head == nil {
			head = n
			return head
		}
		resto = head.prox
		head.prox = n
		n = head
		head = resto
	}

}

func main() {
	head := &ll{val: 0, prox: nil}
	for i := 1; i < 20; i++ {
		add_next(i, head)
	}

	show_list(head)
	head = inverter(head)
	fmt.Printf("\n")
	show_list(head)
}
