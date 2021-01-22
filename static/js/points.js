'use strict';

const points = fetch("/points.json")
    .then(response => response.json());

const NotFoundMessage = "Your username was not found."
const StudentSubmitBtn = document.getElementById("student_username_submit");
const StudentInput = document.getElementById("student_username_input");
const PointsDiv = document.getElementById("student_results");

StudentSubmitBtn.addEventListener("click", (e) => handlePointsQuery());
StudentInput.addEventListener('keypress', function (e) {
    if (e.key === 'Enter') handlePointsQuery();
});

function handlePointsQuery() {
    Promise.resolve(points)
        .then(points => {
            let username = StudentInput.value;

            const shaObj = new jsSHA("SHA-256", "TEXT", { encoding: "UTF8" });
            shaObj.update(username);
            const hash = shaObj.getHash("HEX");

            const students = points["Students"];
            const tasks = points["Tasks"];

            let found = false;
            PointsDiv.innerHTML = "";

            for (let year in students) {
                if (students[year][hash]) {
                    found = true;
                    PointsDiv.appendChild(
                        generateForYear(
                            year,
                            username,
                            tasks[year],
                            students[year][hash]
                        )
                    );
                }
            }

            if (!found) {
                let notFoundP = document.createElement("p");
                notFoundP.innerText = NotFoundMessage;
                PointsDiv.appendChild(notFoundP);
            }
        })
}

function generateForYear(year, username, yearTasks, yearStudent) {
    let yearDiv = document.createElement("div");
    yearDiv.className = "year-points-div";

    let numTasks = yearTasks.length;

    let titleText = "Points for ";
    titleText = titleText.concat(username);
    titleText = titleText.concat(" in year ");
    titleText = titleText.concat(year);

    if (numTasks == 1) {    // showing total points only
        titleText = titleText.concat(": ");

        let pointsText = document.createElement("b");
        pointsText.textContent = yearStudent["TaskPoints"][0];
        pointsText.setAttribute("id", "points_line_total");

        let yearTitle = document.createElement("p");
        yearTitle.setAttribute("id", "points_line_prefix");
        yearTitle.textContent = titleText;

        yearDiv.appendChild(yearTitle);
        yearDiv.appendChild(pointsText);
        return yearDiv;
    }

    let yearTitle = document.createElement("h3");
    yearTitle.textContent = titleText;

    let yearTable = document.createElement("table");
    yearTable.className = "table";

    let tableHeaderRow = document.createElement("tr");
    let taskColumnHead = document.createElement("th");
    taskColumnHead.textContent = "Task";
    let pointsColumnHead = document.createElement("th");
    pointsColumnHead.textContent = "Points Received";

    tableHeaderRow.appendChild(taskColumnHead);
    tableHeaderRow.appendChild(pointsColumnHead);
    yearTable.appendChild(tableHeaderRow);

    for (let i = 0; i < numTasks - 1; i++) {
        let taskRow = document.createElement("tr");

        let taskCell = document.createElement("td");
        taskCell.textContent = yearTasks[i];

        let pointsCell = document.createElement("td");
        pointsCell.textContent = yearStudent["TaskPoints"][i];

        taskRow.appendChild(taskCell);
        taskRow.appendChild(pointsCell);
        yearTable.appendChild(taskRow);
    }

    let totalRow = document.createElement("tr");
    let totalTextCell = document.createElement("th");
    totalTextCell.textContent = "Total Points";
    let totalPointsCell = document.createElement("th");
    totalPointsCell.textContent = yearStudent["TaskPoints"][numTasks - 1];

    totalRow.appendChild(totalTextCell);
    totalRow.appendChild(totalPointsCell);
    yearTable.appendChild(totalRow);

    yearDiv.appendChild(yearTitle);
    yearDiv.appendChild(yearTable);
    return yearDiv;
}