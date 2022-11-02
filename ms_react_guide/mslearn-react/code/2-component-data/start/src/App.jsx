import './index.css'
import React from 'react';
import RecipeTitle from "./RecipeTitle";

// TODO: Import IngredientList
import IngredientList from "./IngredientList";

import Procedure from "./Procedure";

function App() {
  // TODO: Add recipe object
  // todo: this is the data we set in the main app file
  const recipe = {
    title:  'Mashed potatoes',
    feedback:  {
      rating:     4.8,
      reviews:    20
    },
    ingredients: [
      { name: '3 potatoes, cut into 1/2" pieces', prepared: false },
      { name: '4 Tbsp butter', prepared: false },
      { name: '1/8 cup heavy cream', prepared: true },
      { name: 'Salt', prepared: false },
      { name: 'Pepper', prepared: true },
    ],
    procedures: [
      { detail: 'add the sliced potatoes to a pan of salted water' },
      { detail: 'make the pan boil' },
      { detail: 'Cook the potatoes until they can be easily penetrated with a fork, about 15-20 minutes' },
      { detail: 'Drain the potatoes' },
      { detail: 'Add butter, cream, salt and pepper to taste' },
      { detail: 'Mashed potatoes' },
      { detail: 'Whisk well again, adding butter and cream as needed' }
    ],
  };

  return (
    // todo: the data post to the component use `props` way
    <article>
      <h1>Recipe Manager</h1>
      {/* TODO: Add RecipeTitle component */}
      <RecipeTitle title={ recipe.title } feedback={ recipe.feedback } />

      {/* TODO: Add IngredientList component */}
      <IngredientList ingredients={ recipe.ingredients } />

      <Procedure procedures={ recipe.procedures }/>

    </article>
  )
}

export default App;
