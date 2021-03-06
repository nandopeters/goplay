
function iWS( server_service ){
	var conn;
	var service=server_service;
	// $("#log")
	//var	displayDivId = $("<div>");
	var displayDivId = $("#log");
	
	setDisplayDiv( displayDivId );
	
	function nullDisplayFunc () {
		return;
	}

	this.setDisplayFunc = function( func ){
		if (func == null)
			appendLog = nullDisplayFunc;
		else
			appendLog = func;
	};

	this.setDisplayFunc(baseDisplayFunc);
	
	this.start = function( message ){
	    if (window["WebSocket"]) {
			conn = new WebSocket(service);
			this.conn = conn;
			if (!conn){
				appendLog(displayDivId,"unable to connect");
				return;
			}

			addOnClose();
			addOnError();
			addOnMessage();
			addOnOpen(message);
	    }
	    else	
		{
	    	appendLog(displayDivId,$("<div><b>Your browser does not support WebSockets.</b></div>"));
		}
	    /*
	    $("#form").submit(function() {
	    	console.log('submit');
	        if (!conn) {
	        	console.log('conn is NULL');
	            return false;
	        }
	        console.log("msg typed:",$("#msg").val() );
			var pMsg = '{"msgtype":"join_session","payload":{"session_id":"51", "msg":"hei there baby"} }';
			console.log('conn is true.  msg=',pMsg);
			this.conn.send(pMsg);

	        return false
	    });
	    **/
	};  
	
	
	function addOnClose(  ){
		conn.onclose = function (evt) {
			console.log("onlcose",evt);
		};
	};
	
	function addOnOpen( message ){
		conn.onopen = function (evt) {
			console.log("onopen",evt);
			if ( typeof message === 'undefined') 
				; // do nothing
			else
				conn.send(message);
		};
	};
	
	this.Send = function ( message){
		conn.send(message);
	};
	
	this.Close = function ( ){
		conn.close();
	};
	
	
	function addOnError(  ){
		conn.onerror = function (evt) {
			console.log("error",evt);
		};
	};
	
	function addOnMessage (  ) {
		conn.onmessage = function(evt) {
			console.log("onmessage",evt);
			appendLog( displayDivId,evt.data);
        }
	};
	
	function setDisplayDiv( divId) {
		// form is $("#anID") 
		displayDivId = divId;
	};

	
	function baseDisplayFunc(logDiv, msg) {
	    var d = logDiv[0]
	    var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
	    logDiv.append(msg+"<br>");
	    if (doScroll) {
	        d.scrollTop = d.scrollHeight - d.clientHeight;
	    }
	};
	
};


