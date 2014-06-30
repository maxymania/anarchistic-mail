var dgram = require('dgram');
var fs = require('fs');


function Server(t,fn){
    var object = {};
    function load(){
        fs.readFile(fn,{encoding:'utf8'},function(err,data){
			console.log("load",fn);
            if(err)return;
            try{
                object = JSON.parse(data);
            }catch(e){}
        });
    }
    var s = dgram.createSocket(t);
    s.bind(64000);
    s.on('message',function(msg,nfo){
        try{
            var key = JSON.parse(''+msg);
            var result = object[key]||null;
            var buf = new Buffer(JSON.stringify(result),'utf8');
            s.send(buf,0,buf.length,nfo.port,nfo.address);
        }catch(e){}
    });
    fs.watch(fn,load);
    load();
}


module.exports = Server;
