
import React from 'react'

function Procedure(props) {

  const proceduresList = props.procedures.map((procedureItem, index) => {
    return (
      <ol>
        { procedureItem.detail }
      </ol>
    );
  });

  // todo: here must add some HTML tag?? if dont add `<div>` will occur error `Objects are not valid as a React child`
  return (
    <div>
      { proceduresList }
    </div>
  );
}

export default Procedure;


