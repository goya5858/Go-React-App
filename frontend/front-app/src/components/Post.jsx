import axios from 'axios';
import GET_all_items from './GET_all_items';
import GET_one_item from './GET_one_item';

const Post = () => {
    
    const handleROOT = async() => {
        axios( {
            method: "get",
            url: "http://localhost/api"
        } )
        .then( res =>{
            console.log(res);
            alert( res.data );
        } )
    }
    
    const handlePOST_item = async() => {
        axios( {
            method: "post",
            url: "http://localhost/api/item",
            data:   {
                        "id": "3",
                        "item_name": "item_3",
                        "price": 7000,
                        "stock": 500,
                    }
        } )
        .then( res =>{
            console.log(res);
            alert( JSON.stringify(res.data) );
        } )
    }

    const handleDELETE_item_1 = async() => {
        axios( {
            method: "delete",
            url: "http://localhost/api/item/1"
        } )
        .then( res =>{
            console.log(res);
            alert( "DELETE Data" );
        } )
    }

    const handlePUT_item_2 = async() => {
        axios( {
            method: "put",
            url: "http://localhost/api/item/2",
            data:   {
                        "id": "2",
                        "item_name": "item_2_2",
                        "price": 2400,
                        "stock": 500,
                    }
        } )
        .then( res =>{
            console.log(res);
            alert( JSON.stringify(res) );
        } )
    }

    return (
        <div>
            <button onClick={handleROOT}>
                ROOT
            </button>

            <GET_all_items/>
            <GET_one_item/>

            <button onClick={handlePOST_item}>
                POST
            </button>

            <button onClick={handleDELETE_item_1}>
                DELETE_1
            </button>

            <button onClick={handlePUT_item_2}>
                PUT_2
            </button>
        </div>
    )
}

export default Post;