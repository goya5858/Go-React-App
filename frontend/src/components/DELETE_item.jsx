import axios from 'axios';
import { useState } from 'react';

const DELETE_item = () => {
    const [target, setTarget] = useState("1")
    
    const handleDELETE_item = async() => {
        axios( {
            method: "delete",
            url: "http://localhost/api/item/" + target
        } )
        .then( res =>{
            console.log(res);
            alert( "DELETE Data" );
        } )
    }
    
    const hundleChange = (event) => {
        const newValue = event.target.value;
        setTarget(newValue);
    }

    return (
        <div>
            <input type="number" value={target} onChange={hundleChange}/>
            <button onClick={handleDELETE_item}>
                DELETE_item
            </button>
        </div>
    )
}

export default DELETE_item;