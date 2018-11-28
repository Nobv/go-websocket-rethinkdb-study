function connect(path) {
    var ws = new WebSocket("ws://localhost:3000/ws/" + path);

    ws.onmessage = function(e) {
        var data = JSON.parse(e.data);

        if(data.old_val === null && data.new_val !== null) {

            var message = data.new_val;
            $('#message-list').append("" + 
                "<li data-id='" + message.id + "'>" +
                    "<div class='view'>" +
                        "<span>" + message.Text + "</span>" +
                    "</div>" +
                "</li>" + 
            "");
        }

    };

}