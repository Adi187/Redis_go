package main

import (
"fmt"
"redis_go/store"
"time"
)

type MetricsMiddleware struct{
	inner store.Storer
	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int
	getMisses   int
	totalSetLatency time.Duration
	totalGetLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
} 

//how to enable 
func (m *MetricsMiddleware) Set(key, value string) error {
	start:= time.Now()
	err:=m.inner.Set(key, value)
    m.setCalls++
	duration:=time.Since(start)
	m.totalSetLatency += duration
	return err
}

func (m *MetricsMiddleware) Get(key string) (string, bool,error) {
	start:=time.Now()
	m.getCalls++
	val,_,err:= m.inner.Get(key)
	if err!=nil{
		m.getMisses++
		return val,false,err
	}

	m.totalGetLatency += time.Since(start)
    return val,true,nil
}

func (m *MetricsMiddleware) Delete(key string) {
	m.deleteCalls++
	m.inner.Delete(key)
}

func (m *MetricsMiddleware) Len() int{
	m.lenCalls++
	return m.inner.Len()
}
 
func (m *MetricsMiddleware) Keys() []string {
	
	return m.inner.Keys() 
}

func (m *MetricsMiddleware) Report() {
	fmt.Printf("GET calls : %d (missed %d)\n",m.getCalls,m.getMisses)
	fmt.Printf("SET calls : %d\n",m.setCalls)
	fmt.Printf("DEL calls : %d\n", m.deleteCalls)
	fmt.Printf("LEN calls: %d\n", m.lenCalls)

	if m.getCalls > 0{
		avg := m.totalSetLatency  / time.Duration(m.getCalls)
		fmt.Printf("AVG GET latency : %v\n", avg)
	}
	if m.setCalls > 0{
		avg := m.totalGetLatency  / time.Duration(m.setCalls)
		fmt.Printf("AVG SET latency : %v\n", avg)
	}
}

func (m *MetricsMiddleware) Incr(key string) (int, error){
	
	return 0,nil
}





