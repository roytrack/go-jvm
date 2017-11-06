package music

import "testing"

func TestOps(t *testing.T){
	mm :=NewMusicManager()
	if mm ==nil{
		t.Error("NewMusicManger failed.")
	}
	if mm.Len()!=0{
		t.Error("NewMusicManager failed,not empty")
	}
	m0 :=&Music{
		"1","MyHeart woll go along","Celion Dion","Pop","http://abc","MP3"}
		mm.Add(m0)
		if mm.Len()!=1{
			t.Error("MusicManger.Add() failed.")
		}
		m :=mm.Find(m0.Name)
		if m==nil{
			t.Error("MusicManger.Find() failed")
		}
		if m.Id !=m0.Id||m.Artist!=m0.Artist||m.Name!=m0.Name||m.Flow!=m0.Flow||m.Source!=m0.Source||m.Type!=m.Type{
			t.Error("MusicManager.Find() failed. Found item mismatch")
		}
		m,err :=mm.Get(0)
		if err!=nil{
			t.Error("MusicManager.Get() failed .",err)
		}
		m =mm.Remove(0)
		if m==nil||mm.Len()!=0{
			t.Error("MusicManger.Remove() failed")
		}
}
