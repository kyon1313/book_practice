

function removeAuthor(button) {
    // Find the parent author div and remove it
    var authorDiv = button.parentElement;
    authorDiv.remove();
}




function addAuthor() {
  const authorsDiv = document.getElementById("authorsAdd");
  const authorDiv = document.querySelector(".authorAdd").cloneNode(true);

  // Reset the First Name input field
  authorDiv.querySelector('input[name="first_name[]"]').value = '';
  authorDiv.querySelector('input[name="middle_name[]"]').value = '';
  authorDiv.querySelector('input[name="last_name[]"]').value = '';
  authorsDiv.appendChild(authorDiv);
  console.log(authorsDiv)
}

function removeAuthor(button) {
  const authorDiv = button.parentNode;
  const authorsDiv = document.getElementById("authorsAdd");
  if (authorsDiv.childElementCount > 1) {
    authorsDiv.removeChild(authorDiv);
  }
}


function removeUpdate(button) {
  console.log("remove update button has been called");
  const authorDiv = button.parentNode;
  const authorsDiv = document.getElementById("authorsUpdate");
  if (authorsDiv.childElementCount > 1) {
    authorsDiv.removeChild(authorDiv);
  }
}



// function addUpdate() {

//   console.log("update author button has been called");
//   const authorsDiv = document.getElementById("authorsUpdate");
//   const authorDiv = document.querySelector(".authorUpdate").cloneNode(true);

//   // Reset the First Name input field
//   authorDiv.querySelector('input[name="first_name[]"]').value = '';
//   authorDiv.querySelector('input[name="middle_name[]"]').value = '';
//   authorDiv.querySelector('input[name="last_name[]"]').value = '';
//   authorsDiv.appendChild(authorDiv);
//   console.log(authorsDiv);
 
// }




