// $(function(){
//     $("#menu_sidebar").load("menu_sidebar.html");
// });

function header(rootDir){
    $.ajax({
        url: rootDir + "menu_sidebar.html", // 読み込むHTMLファイル
        cache: false,
        async: false,
        dataType: 'html',
        success: function(html){
            html = html.replace(/\{\$root\}/g, rootDir); //header.htmlの{$root}を置換
            document.write(html);
        }
    });
}