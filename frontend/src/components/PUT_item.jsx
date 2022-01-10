import axios from 'axios';
import { useState } from 'react';

const PUT_item = () => {
    const [id, setId] = useState("2");
    const [item_name, setItemName] = useState("item_2_2");
    const [price, setPrice] = useState(2400);
    const [stock, setStock] = useState(500);
    
    const handlePUT_item = async() => {
        axios( {
            method: "put",
            url: "http://localhost/api/item/" + id,
            data:   {
                        "item_name": item_name,
                        "price": Number(price),
                        "stock": Number(stock),
                    }
        } )
        .then( res =>{
            console.log(res);
            alert( JSON.stringify(res) );
        } )
    }
    
    const hundleIdChange = (event) => {
        setId( event.target.value );
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
            <input type="number" value={id} onChange={hundleIdChange}/>
            <input type="text" value={item_name} onChange={hundleNameChange}/>
            <input type="number" value={price} onChange={hundlePriceChange}/>
            <input type="number" value={stock} onChange={hundleStockChange}/>

            <button onClick={handlePUT_item}>
                PUT
            </button>
        </div>
    )
}

export default PUT_item;