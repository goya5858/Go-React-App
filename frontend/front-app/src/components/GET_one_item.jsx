import axios from 'axios';

const handleGET_item_2 = async() => {
    axios( {
        method: "get",
        url: "http://localhost/api/item/2"
    } )
    .then( res =>{
        console.log(res);
        alert( JSON.stringify(res.data) );
    } )
}

const GET_one_item = () => {
    return (
        <div>
            <button onClick={handleGET_item_2}>
                GET_item_2
            </button>
        </div>
    )
}

export default GET_one_item;