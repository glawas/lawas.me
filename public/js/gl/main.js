var captcha = {
    show: function() {
        if (null === document.getElementById('contact-email_format')) {
            var _t = $("#contact-form-action-wrapper");
            _t.prepend($("<label id=\"contact-email_format\" class=\"pull-left;\"><input type=\"checkbox\" value=\"html\" name=\"email_format\" />&nbsp;I am human.</label>"));
        }
    }
};

$(document).ready(function() {
    var _cForm = $('#contact-form');
    if (_cForm.length > 0) {
        _cForm.find("input[type=text],input[type=email],textarea").each(function() {
            $(this).focus(function(e) {
                captcha.show();
            });
        });
        //_cForm.validate();
    }
});
