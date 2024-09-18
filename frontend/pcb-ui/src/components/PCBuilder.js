import React, { useState, useEffect } from 'react';

function PCBuilder() {
  const [components, setComponents] = useState([]);
  const [selectedComponents, setSelectedComponents] = useState({});
  const [budget, setBudget] = useState(1000);

  useEffect(() => {
    fetch('http://localhost:8080/components?budget=' + budget)
      .then((res) => res.json())
      .then((data) => setComponents(data));
  }, [budget]);

  const handleSelect = (category, component) => {
    setSelectedComponents({
      ...selectedComponents,
      [category]: component,
    });
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">PC Builder</h1>
      <label className="block mb-2">
        Budget: 
        <input
          type="number"
          value={budget}
          onChange={(e) => setBudget(e.target.value)}
          className="ml-2 border rounded p-2"
        />
      </label>
      <div className="grid grid-cols-2 gap-4">
        {components.map((component) => (
          <div key={component.id}>
            <h3 className="text-lg">{component.category}</h3>
            <select
              onChange={(e) =>
                handleSelect(component.category, e.target.value)
              }
              className="border p-2 w-full"
            >
              {component.options.map((option) => (
                <option key={option.id} value={option.id}>
                  {option.name} - ${option.price}
                </option>
              ))}
            </select>
          </div>
        ))}
      </div>
    </div>
  );
}

export default PCBuilder;
