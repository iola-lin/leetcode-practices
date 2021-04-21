<?php

$nums = [1,1,2];
$len = removeDuplicate($nums);
var_dump($len);
var_dump($nums);


$nums2 = [0,0,1,1,1,2,2,3,3,4];
$len2 = removeDuplicate($nums2);
var_dump($len2);
var_dump($nums2);


function removeDuplicate(&$input) { 
    $len = 0;
    foreach ($input as $idx => $elem) {
        if ($idx == 0) { $len++; continue; }
        if ($elem !== $input[$len-1]) {
            $input[$len] = $elem;
            $len++;
        }
        unset($input[$idx]);
    }
    return $len;
}