$(document).ready(function () {
        
        $(document).on('click', '#toggle-view li h2', function () {
         //   var text = $('#toggle-view li').children('div.panel'); 
            var text = $(this).parent().children('div.panel');

            if (text.is(':hidden')) {
                text.slideDown('200');
                $(this).parent().children('span').html('-');
            } else {
                text.slideUp('200');
                $(this).parent().children('span').html('+');
            }

    });

});



