var editor;
var info = "请在此填写回复内容...";

KindEditor.ready(function(K) {
	editor = K.create('textarea[name="content"]', {
		afterFocus : function() {focusMessage()},
		afterBlur : function() {infoMessage()},
		resizeType : 1,
		allowFileManager : true,
		allowPreviewEmoticons : true,
		allowImageUpload : true,
		items : [
			'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold', 'italic', 'underline',
			'removeformat', '|', 'justifyleft', 'justifycenter', 'justifyright', 'insertorderedlist',
			'insertunorderedlist', '|', 'emoticons', 'image', 'link', '|', 'baidumap', 'table']
	});
});

function focusMessage() {
	if (editor.html() == info)
	{
		editor.html('');
	}
}


function infoMessage() {

	if ($.trim(editor.html()) == '' || $.trim(editor.html()).length == 0)
	{
		editor.html(info);
	}
}