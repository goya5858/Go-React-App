import axios from 'axios';
import { useState } from 'react';

const POST_item = () => {
    const [item_name, setItemName] = useState("item_3");
    const [price, setPrice] = useState(7000);
    const [stock, setStock] = useState(500);
    
    const handlePOST_item = async() => {
        axios( {
            method: "post",
            url: "http://localhost/api/item",
            data:   {
                        //"id": "3",
                        "item_name": item_name,
                        "price": Number(price), //本来であればTSXで型固定が良い
                        "stock": Number(stock), //本来であればTSXで型固定が良い
                    }
        } )
        .then( res =>{
            console.log(res);
            alert( JSON.stringify(res.data) );
        } )
    }
    
    const hundleNameChange = (event) => {
        setItemName( event.target.value );
    }
    const hundlePriceChange = (event) => {
        setPrice( event.target.value )
    }
    const hundleStockChange = (event) => {
        setStock( event.target.value )
    }

    return (
        <div>
            <input type="text" value={item_name} onChange={hundleNameChange}/>
            <input type="number" value={price} onChange={hundlePriceChange}/>
            <input type="number" value={stock} onChange={hundleStockChange}/>

            <button onClick={handlePOST_item}>
                POST
            </button>
        </div>
    )
}

export default POST_item;