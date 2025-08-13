# Unleash Feature Flag Demo

This repository demonstrates how to implement feature flags using **Unleash** combined with trunk-based development. Feature flags decouple deployment from release, reduce risk, and enable continuous delivery. By using trunk-based development instead of complex Git flows (GitFlow, GitHub Flow), you maintain a linear history, minimize long-lived branches, and avoid merge conflicts—every commit goes directly to `main` with feature guards around unfinished work.

## Demo Video

[![Watch the Demo](unleash.png)](https://www.youtube.com/watch?v=uvuraFixE_Y)

## Tech Stack

- Backend: NestJS
- Frontend: React
- Feature Flag Service: Unleash

## Backend & Frontend Implementation

To explore the backend and frontend implementations in detail, visit the following repositories:

- [Backend Repository](https://github.com/BooRuleDie/FF-Backend)
- [Frontend Repository](https://github.com/BooRuleDie/FF-Frontend)

Here's a overview:

### Frontend

```tsx
export const Admins: React.FC = () => {
    const { admins, isLoading: isLoadingAdmins } = useGetAdmins();
    const isFlagAdminUsersEnabled = useFlag("admins");

    if (!isFlagAdminUsersEnabled) return "";

    return (
        <>
            <h1 className="mb-3 text-xl font-bold">Admins</h1>
            {isLoadingAdmins ? (
                <div className="flex justify-center items-center h-64">
                    <Spin size="large" />
                </div>
            ) : (
                <Table<User>
                    columns={columns}
                    dataSource={admins}
                    className="max-w-200 bg-white rounded-2xl border border-gray-200 shadow-md pt-2 overflow-x-auto"
                    rowKey="id"
                />
            )}
        </>
    );
};
```

### Backend

```ts
@Get('/admins')
getAdmins(): User[] {
    const isFlagEnabled = this.unleash.isEnabled('admins');
    if (isFlagEnabled) {
        return this.appService.getAdmins();
    }
    throw new UnauthorizedException();
}
```

## Upsides

1. Continuous Integration: Fewer merge conflicts and integration problems.
2. Git History: Linear, simpler, and cleaner history (PR + squash & rebase).
3. Deployment: Toggle features on/off instantly—no redeploy needed.
4. A/B Testing: Expose features to specific user segments.
5. Team Productivity: Eliminate long-lived branch conflicts; enable parallel work.

## Downsides

1. Flag Management: Remember to clean up flags once features stabilize.
2. Complex Codebase: Conditional logic can add clutter.
3. Infrastructure Requirements: High reliance on unit & integration tests since every commit updates all environments.
