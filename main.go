package main

import (
	"flag"
	"fmt"
	"time"
	"traitement_images/filter"
	"traitement_images/task"
)

func main() {
	srcdir := flag.String("src", "imgs", "Répertoire d'entrée")
	dstdir := flag.String("dst", "output", "Répertoire de sortie")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgrp", "waitgrp/channel")
	poolsize := flag.Int("poolsize", 4, "Taille du pool de travailleurs pour la tâche de canal.")
	flag.Parse()

	var f task.Filter

	switch *filterType {
	case "grayscale":
		f = &filter.GrayscaleFilter{}
	case "blur":
		f = &filter.BlurFilter{}
	default:
		fmt.Println("Le type du filtre est invalide")
		return
	}

	var t task.Tasker

	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcdir, *dstdir, f)
	case "channel":
		t = task.NewChanTask(*srcdir, *dstdir, f, *poolsize)
	default:
		fmt.Println("Invalid task type")
		return
	}

	start := time.Now()

	err := t.Process()
	if err != nil {
		fmt.Printf("Erreur lors du traitement de l'image: %s\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Le traitement de l'image a pris %s\n", elapsed)
}
