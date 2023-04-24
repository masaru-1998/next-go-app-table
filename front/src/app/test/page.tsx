async function getData() {
    const res = await fetch('https://api.sampleapis.com/codingresources/codingResources');
    // The return value is *not* serialized
    // You can return Date, Map, Set, etc.

    // Recommendation: handle errors
    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data');
    }

    return res.json();
}

export default async function Page() {
    const data = await getData();
    console.log(data[0].id);
    return (
        <p>{ data[0].id }</p>
    );
}