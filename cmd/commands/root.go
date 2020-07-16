package commands

import (
	"github.com/Kifen/enkrypt/pkg/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	source, target, key string
	port int32
	errCh = make(chan error)
	done = make(chan struct{})
	grpcServer *util.EnkryptServer
)

var rootCmd = &cobra.Command{
	Use: "enkrypt",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		validateArgs(source, target)
		grpcServer = util.NewServer(int(port), key, source, target)
		go mirrorSource()
		go grpcServer.Serve()
		waitOsSignals()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "seed phrase")
	rootCmd.Flags().StringVarP(&source, "source", "s", "", "source directory")
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "target directory")
	rootCmd.Flags().Int32VarP(&port, "port", "p", 2000, "server port")
}

// Execute executes root CLI command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func validateArgs(args ...string){
	for _, arg := range args {
		i, err := util.ValidatePath(arg)
		if err != nil {
			log.Fatalf("Invalid arg %s", arg)
		}

		if !i.IsDir() {
			log.Fatalf("%s is not a directory", arg)
		}
	}
}

func mirrorSource(){
	err := util.CopyDir(source, target)
	if err != nil {
		log.Fatal(err)
	}

	util.Done <- struct{}{}
	log.Println("Files copied...")

	f, err := util.EncryptFolder(target, key)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.UpdateEncryptedFolder(f)

	log.Printf("Source directory %s mirrored successfully.", source)
	log.Printf("Target directory %s encrypted", target)
}

func waitOsSignals() {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}...)
	<-ch

	go func() {
		select {
		case s := <-ch:
			log.Fatalf("Received signal %s: terminating", s)
		}
	}()
}