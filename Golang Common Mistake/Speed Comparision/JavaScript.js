var date1=new Date()
var start=date1.getTime()
var limit=1000
var mylist=new Array()
for (var i=0;i<limit;i++){
    mylist[i]=i
    console.log(mylist[i])
}
var date2=new Date()
var end=date2.getTime()
var duration =(end-start)
console.log("Execution Time="+duration+" ms")