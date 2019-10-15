$('#test_href').click(function(){
 	console.log($(this).text());
	$('#secret_block1').css('display','block');
	$('#secret_block2').css('display','block');
})
$('#secret_block1').click(function(){
	$('#secret_block1').css('display','none');
	$('#secret_block2').css('display','none');
})
function fibanachi(){
	var number = parseInt($('#fibonachi_number').val());
	if(!number){
		alert('ошибка некоректное значение')
	}else{
		$.ajax({
			type:'POST',
			url:'fibonachi.php',
			data:{number}, 
			success : function(data){
				$('#result').text(`Результат F(${$('#fibonachi_number').val()})=${data}`);
				console.log(data);
			},
			error : function(error){
				console.log(error);
			}
		});
	}
}
$('#btn1').click(function(){
	fibanachi();
})
$('#fibonachi_number').keyup(function(e){
	if (e.keyCode == 13) {
		fibanachi();
	}
})