package project

import (
	"github.com/golangdaddy/leap/models"
)

func buildStructure(config models.Config) *models.Stack {

	tree := &models.Stack{
		WebsiteName: "Grocery Store Management System",
		Config:      config,
		Options: models.StackOptions{
			ChatGPT: true, // For potential customer service automation and queries
		},
	}

	// Define product categories
	category := &models.Object{
		Context: "Category to which products are assigned (e.g., Dairy, Produce)",
		Parents: []string{},
		Name:    "category",
		Fields: []*models.Field{
			{
				Context:  "The name of the category, such as 'Dairy', 'Produce', or 'Bakery'. This helps in classifying products under relevant sections, making inventory management and shopping easier.",
				Name:     "name",
				JSON:     "string_100",
				Required: true,
			},
			{
				Context:  "A brief description of what types of products fall under this category. For example, the 'Dairy' category might include milk, cheese, and yogurt.",
				Name:     "description",
				JSON:     "string_1000",
				Required: false,
			},
		},
		Options: models.Options{},
	}

	// Define products under each category
	product := &models.Object{
		Context: "Products that are sold in the grocery store",
		Parents: []string{category.Name},
		Name:    "product",
		Fields: []*models.Field{
			{
				Context:  "The specific name of the product, such as 'Organic Almond Milk' or 'Whole Wheat Bread'. This is used for listing and identifying stock in the inventory system.",
				Name:     "name",
				JSON:     "string_100",
				Required: true,
			},
			{
				Context:  "The retail price of the product, which is used in billing and sales analysis.",
				Name:     "price",
				JSON:     "number_float",
				Required: true,
			},
			{
				Context:  "The current stock level of the product, critical for inventory management and reorder processes.",
				Name:     "stock",
				JSON:     "number_int",
				Required: true,
			},
			{
				Context:  "The expiry date of the product, important for managing perishable items and ensuring quality control.",
				Name:     "expiryDate",
				JSON:     "string_date",
				Required: false,
			},
		},
		Options: models.Options{},
	}

	// Define staff members
	staff := &models.Object{
		Context: "Information about staff members working at the grocery store",
		Parents: []string{},
		Name:    "staff",
		Fields: []*models.Field{
			{
				Context:  "The full legal name of the staff member, used for personnel records and payroll.",
				Name:     "fullName",
				JSON:     "string_100",
				Required: true,
			},
			{
				Context:  "The job position or title of the staff member, crucial for defining roles and responsibilities within the store.",
				Name:     "position",
				JSON:     "string_100",
				Required: true,
			},
			{
				Context:  "The date when the staff member was hired, important for tracking employment duration and benefits eligibility.",
				Name:     "hireDate",
				JSON:     "string_date",
				Required: true,
			},
			{
				Context:  "The salary of the staff member, necessary for payroll processing and financial planning.",
				Name:     "salary",
				JSON:     "number_float",
				Required: true,
			},
		},
		Options: models.Options{},
	}

	// Define transactions for purchase history
	transaction := &models.Object{
		Context: "Records of transactions made by customers",
		Parents: []string{},
		Name:    "transaction",
		Fields: []*models.Field{
			{
				Context:  "A unique identifier for the transaction, used for sales tracking and customer service inquiries.",
				Name:     "transactionID",
				JSON:     "string_30",
				Required: true,
			},
			{
				Context:  "The name of the customer involved in the transaction, useful for personalized marketing and customer relationship management.",
				Name:     "customerName",
				JSON:     "string_100",
				Required: false,
			},
			{
				Context:  "The total monetary amount of the transaction, critical for financial reporting and analysis.",
				Name:     "totalAmount",
				JSON:     "number_float",
				Required: true,
			},
			{
				Context:  "The date on which the transaction occurred, important for accounting and historical data analysis.",
				Name:     "date",
				JSON:     "string_date",
				Required: true,
			},
			{
				Context:  "The method of payment used in the transaction, such as cash, credit card, or digital payment, important for reconciling accounts and customer preferences.",
				Name:     "paymentMethod",
				JSON:     "string_50",
				Required: true,
			},
		},
		Options: models.Options{},
	}

	// Add all objects to the tree
	tree.Objects = append(tree.Objects, category, product, staff, transaction)

	return tree
}
