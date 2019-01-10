function changeRace() {
    //Friendly Race Selection
    var x = document.getElementsByName("FRaceSelect");
    var FriendlyRace = x[0][x[0].selectedIndex].text

    list = document.getElementsByClassName("FSpecialties")
    for (let element of list) {
        

        //true if element is of the current race or "any"
        var b = element.classList.contains(FriendlyRace) || element.classList.contains("any");

        if(b){
            element.style.display = 'block'
        }else{
            element.style.display = 'none'
            element.childNodes[1].checked = false;//childnode[1] is the input tag if that changes then this will also need to change this is a hack. To clear race specific checked boxes when you change race.
        }
    }

    //repeated for Enemy
    var x = document.getElementsByName("ERaceSelect");
    var EnemyRace = x[0][x[0].selectedIndex].text

    list = document.getElementsByClassName("ESpecialties")
    for (let element of list) {
        

        //true if element is of the current race or "any"
        var b = element.classList.contains(EnemyRace) || element.classList.contains("any");

        if(b){
            element.style.display = 'block'
        }else{
            element.style.display = 'none'
            element.childNodes[1].checked = false;//childnode[1] is the input tag if that changes then this will also need to change this is a hack.
        }
    }
}