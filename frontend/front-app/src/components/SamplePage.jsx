import axios from 'axios';
import GET_all_items from './GET_all_items';
import GET_one_item from './GET_one_item';
import POST_item from './POST_item';
import DELETE_item from './DELETE_item';
import PUT_item from './PUT_item';

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

    

    return (
        <div>
            <button onClick={handleROOT}>
                ROOT
            </button>

            <GET_all_items/>
            <GET_one_item/>
            <POST_item/>
            <DELETE_item/>
            <PUT_item/>
        </div>
    )
}

export default Post;