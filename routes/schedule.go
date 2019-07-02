package routes

import (
	"github.com/Soontao/PDISolutionCenter/models"
)

// AddNewSchedulePlan endpoint
func AddNewSchedulePlan(c *RouteContext) {

	payload := &models.Schedule{}

	c.HTTP.Bind(payload)

	job, err := c.Services.ScheduleService.AddJob(payload)

	if err != nil {
		c.HTTP.AbortWithStatusJSON(500, PlainJSONObject{"error": err.Error()})
		return
	}

	c.HTTP.JSON(201, job)

}
