// var _cdn = 'http://lawas-static.s3-website-us-east-1.amazonaws.com';
var _cdn = "http://dev-glawas.dvidshub.net:8686"
var date_options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };

$(document).ready(function(e){
    var _url = _cdn + '/js/projects.json';
    var projects_div = $("#projects_wrapper");
    $.get(_url, {}, function(data, status) {
        for (var i = 0; i < data.length; i++) {
            var _d = new Date(data[i].date);
            projects_div.append(
                '<div class="project_item">'
                + '<h4>'+data[i].name+'<br><small class="text-muted">'+_d.toLocaleDateString("en-US", date_options)+'</small></h4>'
                + data[i].description
                + '<div class="row">'
                + '<div class="col-sm-6">'+ data[i].url + '</div>'
                + '<div class="col-sm-6"><span class="pull-right">'+ data[i].tags.join(", ") + '</span></div>'
                + '</div>'
                + '</div>'
            );
        }
    }, 'json');
});