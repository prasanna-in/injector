package injector

import (
	"os"
	"syscall"
)

type Proc struct {
	Process *os.Process
}


func GetProcess(processID int) (Proc, error) {
	process, err := os.FindProcess(processID)
	return Proc{Process:process}, err
}


func (p *Proc ) Attach(f func(bool,  *Proc)(error))(error)  {
	err := syscall.PtraceAttach(p.Process.Pid)
	if err != nil{
		return err
	}
	return f(true,p)

}

func (p *Proc) GetRegisters(f func(*syscall.PtraceRegs)(*syscall.PtraceRegs, error))(*syscall.PtraceRegs, error)  {
	regs := &syscall.PtraceRegs{}
	err := syscall.PtraceGetRegs(p.Process.Pid,regs)
	if err != nil {
		return nil,err
	}
	return f(regs)
}

func (p *Proc) SetRegisters(regs *syscall.PtraceRegs) (error) {
	return syscall.PtraceSetRegs(p.Process.Pid, regs)
}

func (p *Proc) PeekData(addr uintptr, out []byte, f func(int, []byte)(int, error)) (int,  error)  {
	count, err := syscall.PtracePeekData(p.Process.Pid,addr,out)
	if count > 0 {
		return f(count,out)
	}
	return count,err
}

func (p *Proc ) PeekText(addr uintptr, out []byte, f func(int, []byte)(int, error))(int, error)  {
	count, err := syscall.PtracePeekText(p.Process.Pid, addr, out)
	if count >0  {
		return f(count,out)
	}
	return count, err
}

func (p *Proc ) PokeData(addr uintptr, out []byte, f func(int, []byte)(int, error)) (int, error) {
	count, err := syscall.PtracePokeData(p.Process.Pid,addr, out)
	if count >0 {
		return f(count, out)
	}
	return count, err
}

func (p *Proc) Countinue(signal int) error {
	return syscall.PtraceCont(p.Process.Pid,signal)
}

func (p *Proc ) PokeText(addr uintptr, out []byte, f func(int, []byte)(int, error)) (int, error) {
	count, err := syscall.PtracePokeData(p.Process.Pid,addr, out)
	if count >0 {
		return f(count, out)
	}
	return count, err
}

