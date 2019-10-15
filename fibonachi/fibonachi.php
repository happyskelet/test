<?php
	require 'config/config.php';
	$conn = mysqli_connect($host, $user, $password, $database);
	if ($conn->connect_error) {
		die("Connection failed: " . $conn->connect_error);    	
	}
	function fibonachi($number,$conn){
		$result1 = mysqli_query($conn,"SELECT result FROM fibonachi WHERE param ='$number'");		
		if (mysqli_num_rows($result1) > 0){			
			while($row = mysqli_fetch_assoc($result1)) {
				$perem = $row["result"];
			}
			return $perem;
		}else{		
			if($number<3){
				return 1;
			}else{
				$focus = fibonachi($number-1,$conn)+fibonachi($number-2,$conn);
				mysqli_query($conn,"INSERT INTO fibonachi (param, result) VALUES ($number,$focus)");
				return $focus;
			}	
		}		
	}
	
	if (isset($_POST['number'])){
	  $number = intval($_POST['number']);		  
	}else{
	  $number = null;		  
	}		
	$result = mysqli_query($conn,"SELECT result FROM fibonachi WHERE param ='$number'");		
	if (mysqli_num_rows($result) > 0){			
		while($row = mysqli_fetch_assoc($result)) {
           $perem = $row["result"];
        }
		echo $perem;			
	}else{		
		$result = fibonachi($number,$conn);	
		echo $result;	
	}
	
	mysqli_close($conn);
?>