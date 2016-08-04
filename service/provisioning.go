package service

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Provisioning_Hook struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningHookService() Provisioning_Hook {
	return Provisioning_Hook{Session: r}
}

func (r *Provisioning_Hook) CreateObject(templateObject *datatypes.Provisioning_Hook) (resp datatypes.Provisioning_Hook, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) EditObject(templateObject *datatypes.Provisioning_Hook) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) GetObject() (resp datatypes.Provisioning_Hook, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Hook) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) GetHookType() (resp datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Hook_Type struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningHookTypeService() Provisioning_Hook_Type {
	return Provisioning_Hook_Type{Session: r}
}

func (r *Provisioning_Hook_Type) GetAllHookTypes() (resp []datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook_Type) GetObject() (resp datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Classification struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceClassificationService() Provisioning_Maintenance_Classification {
	return Provisioning_Maintenance_Classification{Session: r}
}

func (r *Provisioning_Maintenance_Classification) GetMaintenanceClassification(maintenanceClassificationId *int) (resp []datatypes.Provisioning_Maintenance_Classification, err error) {
	params := []interface{}{
		maintenanceClassificationId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Classification) GetMaintenanceClassificationsByItemCategory() (resp []datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Classification) GetObject() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Classification) GetItemCategories() (resp []datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Classification_Item_Category struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceClassificationItemCategoryService() Provisioning_Maintenance_Classification_Item_Category {
	return Provisioning_Maintenance_Classification_Item_Category{Session: r}
}

func (r *Provisioning_Maintenance_Classification_Item_Category) GetObject() (resp datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Classification_Item_Category) GetMaintenanceClassification() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Slots struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceSlotsService() Provisioning_Maintenance_Slots {
	return Provisioning_Maintenance_Slots{Session: r}
}

func (r *Provisioning_Maintenance_Slots) GetObject() (resp datatypes.Provisioning_Maintenance_Slots, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Ticket struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceTicketService() Provisioning_Maintenance_Ticket {
	return Provisioning_Maintenance_Ticket{Session: r}
}

func (r *Provisioning_Maintenance_Ticket) GetObject() (resp datatypes.Provisioning_Maintenance_Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Ticket) GetAvailableSlots() (resp datatypes.Provisioning_Maintenance_Slots, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Ticket) GetMaintenanceClass() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Ticket) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Window struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceWindowService() Provisioning_Maintenance_Window {
	return Provisioning_Maintenance_Window{Session: r}
}

func (r *Provisioning_Maintenance_Window) AddCustomerUpgradeWindow(customerUpgradeWindow *datatypes.Container_Provisioning_Maintenance_Window) (resp bool, err error) {
	params := []interface{}{
		customerUpgradeWindow,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceClassifications() (resp []datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceStartEndTime(ticketId *int) (resp datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindowForTicket(maintenanceWindowId *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		maintenanceWindowId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindowTicketsByTicketId(ticketId *int) (resp []datatypes.Provisioning_Maintenance_Ticket, err error) {
	params := []interface{}{
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindows(beginDate *time.Time, endDate *time.Time, locationId *int, slotsNeeded *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		beginDate,
		endDate,
		locationId,
		slotsNeeded,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenceWindows(beginDate *time.Time, endDate *time.Time, locationId *int, slotsNeeded *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		beginDate,
		endDate,
		locationId,
		slotsNeeded,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) UpdateCustomerUpgradeWindow(maintenanceStartTime *time.Time, newMaintenanceWindowId *int, ticketId *int) (resp bool, err error) {
	params := []interface{}{
		maintenanceStartTime,
		newMaintenanceWindowId,
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionService() Provisioning_Version1_Transaction {
	return Provisioning_Version1_Transaction{Session: r}
}

func (r *Provisioning_Version1_Transaction) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetLoopback() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetPendingTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetTicketScheduledActionReference() (resp []datatypes.Ticket_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction) GetTransactionStatus() (resp datatypes.Provisioning_Version1_Transaction_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction_Group struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionGroupService() Provisioning_Version1_Transaction_Group {
	return Provisioning_Version1_Transaction_Group{Session: r}
}

func (r *Provisioning_Version1_Transaction_Group) GetAllObjects() (resp []datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction_Group) GetObject() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction_History struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionHistoryService() Provisioning_Version1_Transaction_History {
	return Provisioning_Version1_Transaction_History{Session: r}
}

func (r *Provisioning_Version1_Transaction_History) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction_History) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction_History) GetTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction_History) GetTransactionStatus() (resp datatypes.Provisioning_Version1_Transaction_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction_Status struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionStatusService() Provisioning_Version1_Transaction_Status {
	return Provisioning_Version1_Transaction_Status{Session: r}
}

func (r *Provisioning_Version1_Transaction_Status) GetNonCompletedTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction_SubnetMigration struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionSubnetMigrationService() Provisioning_Version1_Transaction_SubnetMigration {
	return Provisioning_Version1_Transaction_SubnetMigration{Session: r}
}