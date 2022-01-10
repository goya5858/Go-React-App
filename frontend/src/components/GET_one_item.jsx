import axios from 'axios';
import { useState } from 'react';

const GET_one_item = () => {
    const [target, setTarget] = useState("2")
    
    const handleGET_item = async() => {
        console.log(target)
        axios( {
            method: "get",
            url: "http://localhost/api/item/" + target
        } )
        .then( res =>{
            console.log(res);
            alert( JSON.stringify(res.data) );
        } )
    }
    
    const hundleChange = (event) => {
        const newValue = event.target.value;
        setTarget(newValue);
    }

    return (
        <div>
            <input type="number" value={target} onChange={hundleChange}/>
            <button onClick={handleGET_item}>
                GET_item
            </button>
        </div>
    )
}

export default GET_one_item;