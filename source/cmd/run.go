package cmd

import (
    "os"
    "strconv"
    "syscall"
    "os/exec"
	"io/ioutil"
    "path/filepath"
    "github.com/spf13/cobra"
    log "github.com/sirupsen/logrus")

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run predefined chroot directory",
	Run: func(cmd *cobra.Command, args []string) {
        run()
	},
}

var contCmd = &cobra.Command{
	Use:   "container_run",
	Run: func(cmd *cobra.Command, args []string) {
        cont()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
    rootCmd.AddCommand(contCmd)
    contCmd.Hidden = true;
}


func run(){
    log.Info("Running ", os.Args[3:], "as ", os.Getpid())
    cmd := exec.Command("/proc/self/exe", append([]string{"container_run"}, os.Args[2:] ... ) ... )
    cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
        UidMappings: []syscall.SysProcIDMap{
            {
                ContainerID: 0,
                HostID: os.Getuid(),
                Size: 1,
            },
        },
         GidMappings: []syscall.SysProcIDMap{
            {
                ContainerID: 0,
                HostID: os.Getgid(),
                Size: 1,
            },
        },
        Unshareflags: syscall.CLONE_NEWNS,
    }
    cmd.Run()
}

    func cont(){
    log.Info("Running ", os.Args[3:], "as ", os.Getpid())

    contgroup()
    hostname := "tent/" + os.Args[2]
    syscall.Sethostname([]byte(hostname))
    syscall.Chroot("/opt/chroot/ubuntu")
    syscall.Chdir("/")
    syscall.Mount("proc", "proc", "proc", 0, "")

    syscall.Mount("/home/bresilla/temp/sshfs", "/opt/chroot/ubuntu/home/bresilla", "auto", 0, "")
    os.Setenv("TERM", "xterm-color")

    cmd := exec.Command(os.Args[3], os.Args[4:] ... )
    cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
    cmd.Run()

    syscall.Unmount("/proc", 0)
}

func contgroup(){
    cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "tent"), 0755)
	ioutil.WriteFile(filepath.Join(pids, "tent/pids.max"), []byte("20"), 0700)
	// Removes the new cgroup in place after the container exits
	ioutil.WriteFile(filepath.Join(pids, "tent/notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(pids, "tent/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
}
