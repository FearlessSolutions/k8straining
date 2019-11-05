import React, { useState, useEffect }  from 'react';
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";

export default function K8WidgetContainer(props) {

    const [results, setResults] = useState([]);
    const [nestCount, setNestCount] = useState(13);
    const [addA, setAddA] = useState(3);
    const [addB, setAddB] = useState(7);
    const [multiplyA, setMultiplyA] = useState(5);
    const [multiplyB, setMultiplyB] = useState(15);
    const [url, setUrl] = useState('');
    const [count, setCount] = useState(0);


    useEffect(() => {
        const fetchData = () => {
            if (url !== '' && count>0) {
                let newResults = results.slice();
                newResults.unshift(`REQUESTING FROM:${url}`);
                fetch(url, {
                    method: url.indexOf('nest') > -1 ? 'get': 'post'
                }).then(function(response) {
                    return response.text();
                }).then(function(data) {
                    console.log(data);
                    newResults.unshift(`RESULT is: ${data}`);
                    setResults(newResults);
                    sleep(350).then(() => {
                        const newCount = count - 1;
                        setCount(count - 1);
                        if (newCount <= 0) {
                            setUrl('')
                        }
                    })
;                });

            }
        };
        fetchData();
    }, [url, count]);

    function reqAdd() {
        let desiredUrl = `/add?a=${addA}&b=${addB}`;
        setCount(1);
        setUrl(desiredUrl);
    }

    function reqMultiply() {
        let desiredUrl = `/multiply?a=${multiplyA}&b=${multiplyB}`;
        setCount(1);
        setUrl(desiredUrl);
    }

    function reqNest(){
        let desiredUrl = `/nest1`;
        setCount(nestCount);
        setUrl(desiredUrl);
    }



    return (
        <div className='container'>
            <div className="row">
                <div className={"col"}>
                    <h2>ENDPOINTS</h2>
                    <Card>
                        <Card.Header>
                                Nest

                        </Card.Header>
                        <Card.Body>
                            Send
                            <input className={"col-sm-2"} value={nestCount} onChange={(e) =>{setNestCount(e.target.value)}}/>
                            requests to the nest endpoint
                            <br/>
                            <Button onClick={() => reqNest()} variant="primary" size="md" active>
                                Send Requests to Nest
                            </Button>


                        </Card.Body>
                    </Card>
                    <Card>
                        <Card.Header>
                            Add
                        </Card.Header>
                        <Card.Body>
                            send request to the add endpoint to add:
                            <br/>
                            <input className={"col-sm-2"} value={addA} onChange={(e) =>{setAddA(e.target.value)}}/>
                            to
                            <input className={"col-sm-2"} value={addB} onChange={(e) =>{setAddB(e.target.value)}}/>
                            <br/>
                            <Button onClick={()=>reqAdd()}>Send</Button>
                        </Card.Body>
                    </Card>
                    <Card>
                        <Card.Header>
                            Multiply
                        </Card.Header>
                        <Card.Body>
                            send request to the multiply endpoint to multiply:
                            <br/>
                            <input className={"col-sm-2"} value={multiplyA} onChange={(e) =>{setMultiplyA(e.target.value)}}/>
                            by
                            <input className={"col-sm-2"} value={multiplyB} onChange={(e) =>{setMultiplyB(e.target.value)}}/>
                            <br/>
                            <Button onClick={()=>reqMultiply()}>Send</Button>
                        </Card.Body>
                    </Card>
                </div>
                <div className="col">
                    <h2>RESULTS <Button onClick={()=>setResults([])}>Clear</Button></h2>
                    <ul>
                        {results.map((result) => <li>{result}</li>)}
                    </ul>
                </div>
            </div>
        </div>
    )
}

const sleep = (milliseconds) => {
    return new Promise(resolve => setTimeout(resolve, milliseconds))
};


