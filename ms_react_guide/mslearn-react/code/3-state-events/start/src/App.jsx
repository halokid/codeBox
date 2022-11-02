import './index.css'
import React, {useEffect, useState} from 'react';
import RecipeTitle from "./RecipeTitle";
import IngredientList from "./IngredientList";

function App() {
  const initialRecipe = {
    title: 'Mashed potatoes',
    feedback: {
      rating: 4.8,
      reviews: 20
    },
    ingredients: [
      {name: '3 potatoes, cut into 1/2" pieces', prepared: false},
      {name: '4 Tbsp butter', prepared: false},
      {name: '1/8 cup heavy cream', prepared: false},
      {name: 'Salt', prepared: true},
      {name: 'Pepper', prepared: true},
    ]
  };

  // TODO: Create recipe state
  const [recipe, setRecipe] = useState(initialRecipe);

  // TODO: Add new state property
  // todo: prepared 的初始值是 false, setPrepared 只能设置在这里的 `prepare` 的值，因为是
  // todo: 专门针对这个值的状态机
  // todo: 新设置的这个state值的初始化值为 `false`
  const [ prepared, setPrepared ] = useState(false);

  // TODO: Create ingredientClick event listener
  function ingredientClick(index) {
    const updatedRecipe = { ... recipe };
    updatedRecipe.ingredients[index].prepared = !updatedRecipe.ingredients[index].prepared;
    setRecipe(updatedRecipe);
  }

  // TODO: Add the effect hook(效果挂钩), 这个是监听数据对象状态的组件，只要是监听的数据的状态产生改变，就会触发下面设定的逻辑
  useEffect(() => {
    // todo: 每当recipe数据改变的时候， 都会触发这个 setPrepared 检查， 这个检查的逻辑是检查recipe数据
    // todo: 的 ingredients 的所有的 prepared 属性是否都是 true, 假如都市true， 则 const[ prepared ]
    // todo: 的 prepared 的值为 true
    // todo: setPrepared 函数里面的参数 set 之后的数据， 就是 const [ prepared ] 里面的这个 prepare的值
    // todo: 代码使用 setPrepared 来更新 prepared。 它使用 every 方法，该方法根据与所指定的条件匹配的每个项返回一个布尔值。 在本例中，我们将检查是否已准备好每项原料。 如果没有，此方法返回 false
    setPrepared(recipe.ingredients.every(i => i.prepared));
  }, [recipe]);

  return (
    <article>
      <h1>Recipe Manager</h1>

      {/* TODO: Pass recipe metadata to RecipeTitle */}
      {/* 引用RecipeTitle组件，引用的内容是RecipeTitle return的内容 */}
      <RecipeTitle title={recipe.title} feedback={recipe.feedback} />

      {/* TODO: Pass ingredients and event listener to IngredientList */}
      <IngredientList ingredients={recipe.ingredients} onClick={ ingredientClick } />

      {/* TODO: Add the prep work display */}
      { prepared ? <h2>Prep work done!</h2> : <h2>Just keep chopping.</h2>}
    </article>
  )
}

export default App;
