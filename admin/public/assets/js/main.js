Array.prototype.unique = function() {
    var a = this.concat();
    for(var i=0; i<a.length; ++i) {
        for(var j=i+1; j<a.length; ++j) {
            if(a[i] === a[j])
                a.splice(j--, 1);
        }
    }

    return a;
};

// self executing function here
(function() {
   // your page initialization code here
   // the DOM will be available here

 //   var nodes = [
	// 	{
	// 		"id": 18,
	// 		"path": "1.18",
	// 		"created_by": 1,
	// 		"name": "gopher.jpg",
	// 		"node_type": 2,
	// 		"created_date": "2014-10-28T15:50:47.303Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 19,
	// 		"path": "1.19",
	// 		"created_by": 1,
	// 		"name": "postgresql.png",
	// 		"node_type": 2,
	// 		"created_date": "2014-10-28T17:53:37.488Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 23,
	// 		"path": "1.23",
	// 		"created_by": 1,
	// 		"name": "Sample picture folder",
	// 		"node_type": 2,
	// 		"created_date": "2014-11-17T16:57:14.654Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 24,
	// 		"path": "1.23.24",
	// 		"created_by": 1,
	// 		"name": "Goku_SSJ3.jpg",
	// 		"node_type": 2,
	// 		"created_date": "2014-11-17T16:58:57.285Z",
	// 		"parent_id": 23,
	// 		"children": []
	// 	}
	// ];

	// function treeify(nodes) {
	//     var indexed_nodes = {}, tree_roots = [];
	//     for (var i = 0; i < nodes.length; i += 1) {
	//         indexed_nodes[nodes[i].id] = nodes[i];
	//     }
	//     for (var i = 0; i < nodes.length; i += 1) {
	//         var parent_id = nodes[i].parent_id;
	//         if (parent_id === 1) {
	//             tree_roots.push(nodes[i]);
	//         } else {
	//             indexed_nodes[parent_id].children.push(nodes[i]);
	//         }
	//     }
	//     return tree_roots;
	// }
	// alert("remember to remove this from main.js:: this is only temporary")
	// alert(JSON.stringify(treeify(nodes), undefined, "\t"));
})();