# Instructions

Implement the functions to pass the tests.


Reference: [link](https://en.wikipedia.org/wiki/Fold_(higher-order_function))

1. Do not change signature of functions
1. Do not change tests





func Fold is designed to perform a fold operation on a list. It takes three arguments: initialValue (an integer), f (a function that takes two integers and returns an integer), and l (a struct). An empty map m is created to store the results. It checks if the list is empty by comparing the length of l.Values to 1. If it is empty, the initialValue is assigned to the map with l.ID as the key, and the map is returned. If the list is not empty, the first value is extracted, and a new list excluding the first value is created. The function then recursively calls itself with the updated parameters, and the final result is returned as the result of the fold operation.

func FoldChan is similar to the previous one but operates on a channel. It takes three arguments: initialValue (an integer), f (a function that takes two integers and returns an integer), and ch (a channel that receives List values). An empty map m is created to store the results. The function iterates over the channel using the range syntax, receiving List values one by one. For each received List, the function calls the Fold function to perform the fold operation. It checks if the ID of the received List already exists in the map m. If it does, it applies the function f to merge the existing value with the folded value and updates the map accordingly. If the ID does not exist, it simply assigns the folded value to the map. Finally, when there are no more values in the channel, the function returns the map with the folded results for each ID.

func FoldChanX is an extension of the previous one and handles multiple channels concurrently. It takes three arguments: initialValue (an integer), f (a function that takes two integers and returns an integer), and chs (variadic parameter representing multiple channels). An empty map m is created to store the results. It uses mutexes to synchronise access to the map m when updating it concurrently. It also utilises waitgroups to ensure that all goroutines have completed their execution. The function loops over the chs slice to iterate over each channel. For each channel, it starts a new goroutine that performs similar steps as the previous function. However, it uses a mutex to lock and unlock access to the map m when updating its values concurrently. Once all the goroutines have completed, the function returns the final map with the folded results for each ID.