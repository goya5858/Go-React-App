import axios from 'axios';

const Post = () => {
    
    const handleROOT = async() => {
        axios( {
            method: "get",
            url: "http://localhost/api"
        } )
        .then( res =>{
            console.log(res)
        } )
    }

    const handleGET_items = async() => {
        axios( {
            method: "get",
            url: "http://localhost/api/items"
        } )
        .then( res =>{
            console.log(res)
        } )
    }

    const handleGET_item_2 = async() => {
        axios( {
            method: "get",
            url: "http://localhost/api/item/2"
        } )
        .then( res =>{
            console.log(res)
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
                        "created_at": "2021-12-28T15:43:23.4346004+09:00",
                        "updated_at": "2021-12-28T15:43:23.4346005+09:00",
                        "deleted_at": "2021-12-28T15:43:23.4346006+09:00"
                    }
        } )
        .then( res =>{
            console.log(res)
        } )
    }

    const handleDELETE_item_1 = async() => {
        axios( {
            method: "delete",
            url: "http://localhost/api/item/1"
        } )
        .then( res =>{
            console.log(res)
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
                        "created_at": "2021-12-28T15:43:23.4346004+09:00",
                        "updated_at": "2021-12-28T15:43:23.4346005+09:00",
                        "deleted_at": "2021-12-28T15:43:23.4346006+09:00"
                    }
        } )
        .then( res =>{
            console.log(res)
        } )
    }

    return (
        <div>
            <button onClick={handleROOT}>
                ROOT
            </button>

            <button onClick={handleGET_items}>
                GET_items
            </button>

            <button onClick={handleGET_item_2}>
                GET_item_2
            </button>

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