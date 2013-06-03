
function iWS( service ){
	var conn;
	
	start(service);
	
	this.func2 = function () {
		
	};
	
	function start( service){
	    if (window["WebSocket"]) {
			conn = new WebSocket(service);
			if (!conn){
				appendLog($("#log"),"unable to connect");
				return;
			}
			console.log(conn);
			addOnClose();
			addOnError();
			addOnMessage();
			addOnOpen();
			addReadyStateChange ();

	    }
	    else	
		{
			appendLog($("#log"),$("<div><b>Your browser does not support WebSockets.</b></div>"));
		}
	    
	    $("#form").submit(function() {
	    	console.log('submit');
	        if (!conn) {
	            return false;
	        }

			var pMsg = '{"msgtype":"join_session","payload":{"schedule_id":"51"} }';
			conn.send(pMsg);

	        return false
	    });
	    
	   
	};  
	
	function addOnClose(  ){
		conn.onclose = function (evt) {
			console.log("onlcose",evt);
			appendLog($("#log"),"connection closed");
		};
	};
	
	function addOnOpen(  ){
		conn.onopen = function (evt) {
			console.log("opopen",evt);
			console.log(conn);
			appendLog($("#log"),"connection Opened");
			var pMsg = '{"msgtype":"join_session","payload":{"schedule_id":"51"} }';
			//conn.send(pMsg);
		};
	};
	
	function addOnError(  ){
		conn.onerror = function (evt) {
			console.log("onerror",evt);
			appendLog($("#log"),"Error");
            appendLog($("#log"),evt.data);
		};
	};
	function addOnMessage (  ) {
		conn.onmessage = function(evt) {
			console.log("onmessage",evt);
            appendLog($("#log"),evt.data);
        }
	};

	function addReadyStateChange (  ) {
		conn.readystatechange = function(evt) {
			console.log("readystatechange",evt);
            appendLog($("#log"),evt.data);
        }
	};
	
	
	function send (msg ){
		conn.send(msg);
	};
	this.func1 = function () {
		
	}
	
};


function appendLog(logDiv, msg) {
    var d = logDiv[0]
    var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
    logDiv.append(msg+"<br>");
    if (doScroll) {
        d.scrollTop = d.scrollHeight - d.clientHeight;
    }
};