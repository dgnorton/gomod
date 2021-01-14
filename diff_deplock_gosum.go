package gomod

type LockSumDiff struct {
	ProjectName string
	DepVer      string
	ModVer      string
}

func DiffLockSum(depLockPath, goSumPath string) ([]*LockSumDiff, error) {
	lock, err := LoadDepLock(depLockPath)
	if err != nil {
		return nil, err
	}

	sum, err := LoadGoSum(goSumPath)
	if err != nil {
		return nil, err
	}

	deps := map[string]string{}
	mods := map[string]string{}
	diffs := []*LockSumDiff{}

	for _, p := range lock.Projects {
		deps[p.Name] = p.Version
	}

	for m, _ := range sum.Modules {
		mods[m.Path] = m.Version

		if depver, ok := deps[m.Path]; !ok {
			diffs = append(diffs, &LockSumDiff{
				ProjectName: m.Path,
				ModVer:      m.Version,
			})
		} else if m.Version != depver {
			diffs = append(diffs, &LockSumDiff{
				ProjectName: m.Path,
				DepVer:      depver,
				ModVer:      m.Version,
			})
		}
	}

	for _, p := range lock.Projects {
		if modver, ok := mods[p.Name]; !ok {
			diffs = append(diffs, &LockSumDiff{
				ProjectName: p.Name,
				DepVer:      p.Version,
			})
		} else if modver != p.Version {
			diffs = append(diffs, &LockSumDiff{
				ProjectName: p.Name,
				DepVer:      p.Version,
				ModVer:      modver,
			})
		}
	}

	return diffs, nil
}
